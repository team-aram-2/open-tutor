import { PUBLIC_API_HOST } from '$env/static/public';
import type { SkillsItem } from '$lib/types/types';

// Get available skills from backend
export async function fetchSkillsJSON<T>(): Promise<T> {
	console.log('starting fetchSkills');
	let response: Response;
	try {
		response = await fetch(PUBLIC_API_HOST + '/skills', {
			method: 'GET',
			credentials: 'include'
		});
		if (!response.ok) {
			console.error(`HTTP error, status: ${response.status}`);
			throw new Error(`HTTP error, status: ${response.status}`);
		}
		console.log(response);
		return (await response.json()) as T;
	} catch (err) {
		console.error('Error fetching skills, ' + err);
	}
	return null as T;
}

export async function loadSkills(): Promise<SkillsItem[] | null> {
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
