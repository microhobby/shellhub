import http from '@/helpers/http';

export const
  setFirewalls = async () => {
    return http().get('/firewall/rules'); 
  };