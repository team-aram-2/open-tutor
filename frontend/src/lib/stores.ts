import { writable, type Writable } from 'svelte/store';

/**
 * Creates persistent svelte store by leveraging localStorage.
 * Does not create localStorage value until server-side rendering is complete.
 */
function createPersistentStore<T>(key: string, initialValue: T): Writable<T> {
	// Attempt to retrieve previously stored value, check for server side rendering
	let storedValue: T | null = null;
	if (typeof window !== 'undefined') {
		const localStorageValue = localStorage.getItem(key);
		storedValue = localStorageValue ? JSON.parse(localStorageValue) : null;
	}
	// Create store
	const store = writable(storedValue !== null ? storedValue : initialValue);

	// Subscribe to store changes to update localStorage
	if (typeof window !== 'undefined') {
		store.subscribe((value) => {
			localStorage.setItem(key, JSON.stringify(value));
		});
	}

	return store;
}
const logged_in = writable(false);
const user_id = writable('');

export { logged_in, user_id };

// Regular stores
export const sessionToken = writable<string | null>(null);

// Persistent stores
export const fontSize = createPersistentStore<string>('fontSize', '16px');
