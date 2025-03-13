import { writable, type Writable } from 'svelte/store';

// Creates persistent svelte store through localStorage
function createPersistentStore<T>(key: string, initialValue:T): Writable<T> {
  // Attempt to retrieve previously stored value
  const storedValue = localStorage.getItem(key);
  // Create store
  const store = writable(storedValue ? JSON.parse(storedValue) : initialValue)

  // Subscribe to store changes to update localStorage
  store.subscribe((value) => {
    localStorage.setItem(key, JSON.stringify(value));
  });

  return store;
}

// Regular stores
export const sessionToken = writable<string | null>(null);

// Persistent stores
export const fontSize = createPersistentStore<number>("fontSize", 16);