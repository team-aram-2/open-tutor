import { logged_in, user_id, user_role_mask } from '$lib/stores';
import { get } from 'svelte/store';
import { jwtDecode } from 'jwt-decode';

export const autoLogin = async () => {
	// Look for the session_token cookie
	const token = getCookie('session_token');

	if (!token) {
		console.log('No session_token cookie found.');
		logged_in.set(false);
		return;
	}

	try {
		const decoded = jwtDecode(token) as { user_id: string; role_mask: number };

		user_id.set(decoded.user_id);
		user_role_mask.set(decoded.role_mask);
		logged_in.set(true);

		console.log('User ID:', get(user_id));
	} catch (e) {
		console.error('Invalid JWT in session_token:', e);
		logged_in.set(false);
	}
};

// Utility: Safe cookie retrieval
function getCookie(name: string): string | null {
	const cookies = document.cookie.split(';');
	for (const cookie of cookies) {
		const [key, value] = cookie.trim().split('=');
		if (key === name) {
			return decodeURIComponent(value);
		}
	}
	return null;
}

export enum Role {
	User = 1 << 0, // 1
	Tutor = 1 << 1, // 2
	Moderator = 1 << 2, // 4
	Admin = 1 << 3 // 8
}

export function getRolesFromMask(mask: number): Role[] {
	return (Object.values(Role) as Role[]).filter(
		(role) => typeof role === 'number' && (mask & role) === role
	);
}

export function getRoleLabel(role: Role): string {
	switch (role) {
		case Role.User:
			return 'User';
		case Role.Tutor:
			return 'Tutor';
		case Role.Moderator:
			return 'Moderator';
		case Role.Admin:
			return 'Admin';
		default:
			return 'Unknown';
	}
}

export function getRoleLabelsFromMask(mask: number): string[] {
	return getRolesFromMask(mask).map(getRoleLabel);
}

export function addRole(mask: number, role: Role): number {
	return mask | role;
}

export function removeRole(mask: number, role: Role): number {
	return mask & ~role;
}

export function hasRole(mask: number, role: Role): boolean {
	return (mask & role) === role;
}
