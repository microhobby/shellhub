<template>
  <div>
    <ListFirewall />
  </div>
</template>

<script>
import ListFirewall from '@/components/firewall/FirewallList.vue';

export default {
  name: 'Firewalls',

  components: {
    ListFirewall
  }
};
</script>