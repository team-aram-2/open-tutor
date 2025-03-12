import { logged_in,user_id } from "$lib/stores";
import { get } from 'svelte/store'
import {jwtDecode} from 'jwt-decode';

export const autoLogin = async () => {
  const cookies = document.cookie.split(";");
  let token = "";

  if (cookies.length === 0 || (cookies.length === 1 && cookies[0] === '')) {
    console.log('No cookies found.');
    return;
  }
  cookies.forEach(cookie => {
    const [, value] = cookie.trim().split('=');
    if(value.includes("Bearer ")) {
      token = value.replace("Bearer ", "");
    }
  });
  if (token === "") {
    logged_in.set(false)
  }
  else {
    logged_in.set(true)
  }
  const decoded = jwtDecode(token) as { user_id: string };
  user_id.set(decoded.user_id);
  console.log(get(user_id))
}