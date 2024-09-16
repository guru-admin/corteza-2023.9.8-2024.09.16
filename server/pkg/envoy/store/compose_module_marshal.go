package store

import (
	"context"
	"strconv"
	"strings"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/service"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/revisions"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/compose/types"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/envoy/resource"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/store"
	systemTypes "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/system/types"
)

func NewComposeModuleFromResource(res *resource.ComposeModule, cfg *EncoderConfig) resourceState {
	return &composeModule{
		cfg: mergeConfig(cfg, res.Config()),

		res: res,

		recFields:  make(map[string]uint64),
		userFields: make(map[string]uint64),
	}
}

func (n *composeModule) Prepare(ctx context.Context, pl *payload) (err error) {
	// Reset old identifiers
	n.res.Res.ID = 0
	n.res.Res.NamespaceID = 0
	for _, rf := range n.res.ResFields {
		rf.Res.ID = 0
		rf.Res.ModuleID = 0
		rf.Res.NamespaceID = 0
	}

	// Get related namespace
	n.relNS, err = findComposeNamespace(ctx, pl.s, pl.state.ParentResources, n.res.RefNs.Identifiers)
	if err != nil {
		return err
	}
	if n.relNS == nil {
		return resource.ComposeNamespaceErrUnresolved(n.res.RefNs.Identifiers)
	}

	n.res.Res.NamespaceID = n.relNS.ID

	// Get related record field modules
	for _, refMod := range n.res.RefMods {
		var mod *types.Module
		if n.relNS.ID > 0 {
			mod, err = findComposeModuleStore(ctx, pl.s, n.relNS.ID, makeGenericFilter(refMod.Identifiers))
			if err != nil {
				return err
			}
		}
		if mod == nil {
			mod = resource.FindComposeModule(pl.state.ParentResources, refMod.Identifiers)
		}
		if mod == nil {
			return composeModuleErrUnresolvedRecordField(refMod.Identifiers)
		}

		for _, i := range refMod.Identifiers {
			n.recFields[i] = mod.ID
		}
	}

	// Get related user field roles
	for _, refRole := range n.res.RefRoles {
		var rl *systemTypes.Role
		rl, err = findRoleStore(ctx, pl.s, makeGenericFilter(refRole.Identifiers))
		if err != nil {
			return err
		}
		if rl == nil {
			rl = resource.FindRole(pl.state.ParentResources, refRole.Identifiers)
		}
		if rl == nil {
			return composeModuleErrUnresolvedUserField(refRole.Identifiers)
		}

		for _, i := range refRole.Identifiers {
			n.userFields[i] = rl.ID
		}
	}

	// Can't do anything else, since the NS doesn't yet exist
	if n.relNS.ID <= 0 {
		return nil
	}

	// Try to get the original module
	n.mod, err = findComposeModuleStore(ctx, pl.s, n.relNS.ID, makeGenericFilter(n.res.Identifiers()))
	if err != nil {
		return err
	}

	// Nothing else to do
	if n.mod == nil {
		return nil
	}

	// Get the original module fields
	// These are used later for some merging logic
	n.mod.Fields, err = findComposeModuleFieldsStore(ctx, pl.s, n.mod)
	if err != nil {
		return err
	}

	if n.mod != nil {
		n.res.Res.ID = n.mod.ID
		n.res.Res.NamespaceID = n.mod.NamespaceID
	}
	return nil
}

func (n *composeModule) Encode(ctx context.Context, pl *payload) (err error) {
	res := n.res.Res
	exists := n.mod != nil && n.mod.ID > 0

	// Determine the ID
	if res.ID <= 0 && exists {
		res.ID = n.mod.ID
	}
	if res.ID <= 0 {
		res.ID = NextID()
	}

	ns := n.relNS

	res.NamespaceID = ns.ID
	if res.NamespaceID <= 0 {
		ns = resource.FindComposeNamespace(pl.state.ParentResources, n.res.RefNs.Identifiers)
		res.NamespaceID = ns.ID
	}
	if res.NamespaceID <= 0 {
		return resource.ComposeNamespaceErrUnresolved(n.res.RefNs.Identifiers)
	}

	if pl.state.Conflicting {
		return nil
	}

	ts := n.res.Timestamps()
	if ts != nil {
		if ts.CreatedAt != nil {
			res.CreatedAt = *ts.CreatedAt.T
		} else {
			res.CreatedAt = *now()
		}
		if ts.UpdatedAt != nil {
			res.UpdatedAt = ts.UpdatedAt.T
		}
		if ts.DeletedAt != nil {
			res.DeletedAt = ts.DeletedAt.T
		}
	}

	// Fields
	var originalFields types.ModuleFieldSet
	if n.mod != nil && n.mod.Fields != nil {
		originalFields = n.mod.Fields
	} else {
		originalFields = make(types.ModuleFieldSet, 0)
	}

	// Get max validatorID for later use
	vvID := make([]uint64, len(res.Fields))
	for i, f := range res.Fields {
		for _, v := range f.Expressions.Validators {
			if vvID[i] < v.ValidatorID {
				vvID[i] = v.ValidatorID
			}
		}
	}

	for i, f := range res.Fields {
		of := originalFields.FindByName(f.Name)
		if of != nil {
			f.ID = of.ID
		} else {
			f.ID = NextID()
		}
		f.ModuleID = res.ID
		f.Place = i
		f.DeletedAt = nil
		f.CreatedAt = *now()

		// Assure validatorIDs
		for j, v := range f.Expressions.Validators {
			if v.ValidatorID == 0 {
				vvID[i] += 1
				v.ValidatorID = vvID[i]

				f.Expressions.Validators[j] = v
			}
		}

		if f.Options != nil && f.Kind == "Record" {
			refMod := f.Options.String("module")
			if refMod == "" {
				refMod = f.Options.String("moduleID")
			}
			modID := n.recFields[refMod]
			if modID <= 0 {
				ii := resource.MakeIdentifiers(refMod)
				mod := resource.FindComposeModule(pl.state.ParentResources, ii)
				if mod == nil || mod.ID <= 0 {
					return composeModuleErrUnresolvedRecordField(ii)
				}
				modID = mod.ID
			}

			f.Options["moduleID"] = strconv.FormatUint(modID, 10)
			delete(f.Options, "module")
		}

		if f.Options != nil && f.Kind == "User" {
			roles := resource.ComposeModuleFieldExtractUserFieldRoles(f.Options["roles"])
			if len(roles) == 0 {
				roles = resource.ComposeModuleFieldExtractUserFieldRoles(f.Options["role"])
			}
			if len(roles) == 0 {
				roles = resource.ComposeModuleFieldExtractUserFieldRoles(f.Options["roleID"])
			}

			var out []string
			for _, r := range roles {
				roleID := n.userFields[r]
				if roleID <= 0 {
					ii := resource.MakeIdentifiers(r)
					role := resource.FindRole(pl.state.ParentResources, ii)
					if role == nil || role.ID == 0 {
						return composeModuleErrUnresolvedUserField(ii)
					}
					roleID = role.ID
				}

				out = append(out, strconv.FormatUint(roleID, 10))
			}

			f.Options["roles"] = out
			delete(f.Options, "role")
			delete(f.Options, "roleID")
		}
	}

	// Evaluate the resource skip expression
	// @todo expand available parameters; similar implementation to compose/types/record@Dict
	if skip, err := basicSkipEval(ctx, n.cfg, !exists); err != nil {
		return err
	} else if skip {
		return nil
	}

	// Create a fresh module
	if !exists {
		err = store.CreateComposeModule(ctx, pl.s, res)
		if err != nil {
			return err
		}

		err = store.CreateComposeModuleField(ctx, pl.s, res.Fields...)
		if err != nil {
			return err
		}
	} else {
		// Update existing module
		switch n.cfg.OnExisting {
		case resource.Skip:
			return nil

		case resource.MergeLeft:
			res = mergeComposeModule(n.mod, res)
			res.Fields = mergeComposeModuleFields(n.mod.Fields, res.Fields)

		case resource.MergeRight:
			res = mergeComposeModule(res, n.mod)
			res.Fields = mergeComposeModuleFields(res.Fields, n.mod.Fields)
		}

		err = store.UpdateComposeModule(ctx, pl.s, res)
		if err != nil {
			return err
		}

		err = store.DeleteComposeModuleField(ctx, pl.s, n.mod.Fields...)
		if err != nil {
			return err
		}
		err = store.CreateComposeModuleField(ctx, pl.s, res.Fields...)
		if err != nil {
			return err
		}

		n.res.Res = res
	}

	var (
		model *dal.Model
		con   = pl.dal.GetConnectionByID(0)
	)

	// convert module to model and assume compose_record for default ident
	if model, err = service.ModuleToModel(ns, res, "compose_record"); err != nil {
		return
	}
	model.ConnectionID = con.ID

	// @note code copied from the service/module.go

	// Set base constraints
	if model.Ident == "compose_record" {
		model.Constraints = map[string][]any{
			"moduleID":    {res.ID},
			"namespaceID": {res.NamespaceID},
		}
	}

	// replace all partition replacement pairs
	rpl := strings.NewReplacer("{{module}}", res.Handle, "{{namespace}}", ns.Slug)
	model.Ident = rpl.Replace(model.Ident)

	// @todo validate ident with connection's ident validator

	// Create the revisions model if enabled
	if res.Config.RecordRevisions.Enabled {
		rModel := revisions.Model()

		// reuse the connection from the module
		rModel.ConnectionID = con.ID
		rModel.Resource = model.Resource
		rModel.ResourceID = NextID()

		if rModel.Ident = res.Config.RecordRevisions.Ident; rModel.Ident == "" {
			rModel.Ident = "compose_record_revisions"
		}

		rModel.Ident = rpl.Replace(rModel.Ident)

		// @todo validate ident with connection's ident validator

		if _, err = pl.dal.ReplaceModel(ctx, nil, rModel); err != nil {
			return
		}
	}

	if _, err = pl.dal.ReplaceModel(ctx, nil, model); err != nil {
		return
	}

	return nil
}
