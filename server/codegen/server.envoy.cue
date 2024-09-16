package codegen

import (
  "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/app"
  "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/codegen/schema"
)


[...schema.#codegen] &
[
  for cmp in app.corteza.components if !cmp.envoy.omit {
    template: "gocode/envoy/yaml_decode.go.tpl"
    output:   "\(cmp.ident)/envoy/yaml_decode.gen.go"
    payload: {
      package: "envoy"

      imports: [
	      "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/\(cmp.ident)/types"
      ]

      componentIdent: cmp.ident
      resources: [ for res in cmp.resources { res }]
    }
  },
  for cmp in app.corteza.components if !cmp.envoy.omit {
    template: "gocode/envoy/store_decode.go.tpl"
    output:   "\(cmp.ident)/envoy/store_decode.gen.go"
    payload: {
      package: "envoy"

      imports: [
	      "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/\(cmp.ident)/types"
      ]

      componentIdent: cmp.ident
      resources: [ for res in cmp.resources { res }]
    }
  },
  for cmp in app.corteza.components if !cmp.envoy.omit {
    template: "gocode/envoy/store_encode.go.tpl"
    output:   "\(cmp.ident)/envoy/store_encode.gen.go"
    payload: {
      package: "envoy"

      imports: [
	      "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/\(cmp.ident)/types"
      ]

      componentIdent: cmp.ident
      resources: [ for res in cmp.resources { res }]
    }
  },
  for cmp in app.corteza.components if !cmp.envoy.omit {
    template: "gocode/envoy/yaml_encode.go.tpl"
    output:   "\(cmp.ident)/envoy/yaml_encode.gen.go"
    payload: {
      package: "envoy"

      imports: [
	      "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/\(cmp.ident)/types"
      ]

      componentIdent: cmp.ident
      resources: [ for res in cmp.resources { res }]
    }
  },


  {
    template: "gocode/envoy/util_rbac.go.tpl"
    output:   "pkg/envoyx/util_rbac.gen.go"
    payload: {
      package: "envoyx"

      components: [for cmp in app.corteza.components if !cmp.envoy.omit {
        ident: cmp.ident,
        resources: cmp.resources
      }]
    }
  },

  {
    template: "gocode/envoy/utils.go.tpl"
    output:   "pkg/envoyx/utils.gen.go"
    payload: {
      package: "envoyx"

      components: [for cmp in app.corteza.components if !cmp.envoy.omit {
        ident: cmp.ident,
        resources: cmp.resources
      }]
    }
  },
]
