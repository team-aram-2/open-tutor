import { PUBLIC_API_HOST } from '$env/static/public';
import type { SkillsItem } from '$lib/types/types';

// Get available skills from backend
export async function fetchSkillsJSON<T>(): Promise<T> {
	console.log('starting fetchSkills');
	const response = await fetch(`${PUBLIC_API_HOST}/skill`, {
		method: 'GET',
		credentials: 'include'
	});
	if (!response.ok) {
		console.error(`HTTP error, status: ${response.status}`);
	}
	return (await response.json()) as T;
}

export async function loadSkills(): Promise<SkillsItem[]> {
	console.log('starting loadSkills');
	const returnedSkills = await fetchSkillsJSON<SkillsItem[]>();
	console.log('returnedSkills: ' + String(returnedSkills));
	return returnedSkills;
}

// On search confirmation (svelte event dispatcher, sends selectedSkills, minRating, pageIndex, pageSize, order, sort to parent component)
export function handleSubmit(event: Event, selectedSkills: string[]) {
	event.preventDefault();
	console.log('Skills selected: ' + String(selectedSkills));

	// use fetch() request to send GET request to backend
	// use Svelte event dispatcher to send information to parent component
}
