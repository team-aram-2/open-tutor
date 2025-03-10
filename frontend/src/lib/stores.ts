import { writable } from 'svelte/store';

export const sessionToken = writable<string | null>(null);