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
const sidebar_width = writable(0)	;

// Regular stores
export const sessionToken = writable<string | null>(null);
export { logged_in, user_id };
export { sidebar_width };

// Persistent stores
export const font_size = createPersistentStore<string>('font_size', '16px');
