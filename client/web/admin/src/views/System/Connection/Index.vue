<template>
  <b-container
    fluid="xl"
    class="d-flex flex-column flex-fill pt-2 pb-3"
  >
    <c-content-header
      :title="$t('title')"
    />

    <div
      class="d-flex flex-column h-100"
    >
      <c-connection-primary
        class="mb-4"
      />

      <c-external-connection-list
        class="flex-fill"
      />

      <!-- <c-federated-node-list /> -->
    </div>
  </b-container>
</template>
<script>
import CConnectionPrimary from 'corteza-webapp-admin/src/components/Connection/CConnectionPrimary'
import CExternalConnectionList from 'corteza-webapp-admin/src/components/Connection/CExternalConnectionList'
// import CFederatedNodeList from 'corteza-webapp-admin/src/components/Connection/CFederatedNodeList'

export default {
  name: 'Connection',

  components: {
    CConnectionPrimary,
    CExternalConnectionList,
    // CFederatedNodeList,
  },

  i18nOptions: {
    namespaces: 'system.connections',
    keyPrefix: 'list',
  },
}
</script>
