import Vue from 'vue';
import { setFirewalls } from '@/api/firewalls';

export default {
  namespaced: true,

  state: {
    firewalls: [],
    numberFirewalls: 0,
  },

  getters: {
    list: (state) => state.firewalls,
    getNumberSessions: (state) => state.numberFirewalls
  },

  mutations: {
    setFirewalls: (state, res) => {

      Vue.set(state, 'firewalls', res.data);
    },
  },

  actions: {
    fetch: async (context) => {
      let res = await setFirewalls();
      context.commit('setFirewalls', res);
    },
  }
};