<template>
  <fragment>
    <div class="d-flex pa-0 align-center">
      <h1>Firewalls</h1>
      <v-spacer />
      <v-spacer />
    </div>

    <v-card class="mt-2">
      <v-app-bar
        flat
        color="transparent"
      >
        <v-toolbar-title />
      </v-app-bar>

      <v-divider />

      <v-card-text class="pa-0">
        <v-data-table
          :headers="headers"
          :items="listFirewalls"
          item-key="uid"
          :sort-by="['started_at']"
          :sort-desc="[true]"
          :items-per-page="10"
          :footer-props="{'items-per-page-options': [10, 25, 50, 100]}"
          :server-items-length="numberFirewalls"
          :options.sync="pagination"
          :disable-sort="true"
        >
          <template v-slot:item.active="{ item }">
            <v-icon
              v-if="item.active"
              color="success"
            >
              check_circle
            </v-icon>
            <v-tooltip
              v-else
              bottom
            >
              <template #activator="{ on }">
                <v-icon v-on="on">
                  check_circle
                </v-icon>
              </template>
              <span>active</span>
            </v-tooltip>
          </template>

          <template v-slot:item.device="{ item }">
            {{ item.action }}
          </template>

          <template v-slot:item.device="{ item }">
            {{ item.user }}
          </template>

          <template v-slot:item.device="{ item }">
            {{ item.source_ip }}
          </template>

          <template v-slot:item.device="{ item }">
            {{ item.hostname }}
          </template>

          <template v-slot:item.device="{ item }">
            {{ item.tenant_id }}
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>
  </fragment>
</template>

<script>
export default {
  name: 'FirewallList',

  data() {
    return {
      numberFirewalls: 0,
      listFirewalls: [],
      pagination: {},

      headers: [
        {
          text: 'Active',
          value: 'active',
          align: 'center'
        },
        {
          text: 'Action',
          value: 'action',
          align: 'center'
        },
        {
          text: 'User',
          value: 'user',
          align: 'center'
        },
        {
          text: 'Source Ip',
          value: 'source_ip',
          align: 'center'
        },
        {
          text: 'Hostname',
          value: 'hostname',
          align: 'center'
        },
        {
          text: 'Tenant Id',
          value: 'tenant_id',
          align: 'center'
        },
        {
          text: 'Actions',
          value: 'actions',
          align: 'center'
        }
      ]
    };
  },

  watch: {
    pagination: {
      async handler () {
        await this.$store.dispatch('firewalls/fetch');
        this.listFirewalls = this.$store.getters['firewalls/list'];
        this.numberFirewalls = this.$store.getters['firewalls/getNumberFirewalls'];
      },
      deep: true
    },
  },
};
</script>
