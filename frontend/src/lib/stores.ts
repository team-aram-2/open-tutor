import { writable } from "svelte/store";

const logged_in = writable(false);
const user_id = writable("");

export {logged_in, user_id}