import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const { id } = params;
	const users = await import('$lib/mock/my_tutors_mock.json');
	const user = users.default.tutors.find((u) => u.userId === id);

	if (!user) {
		throw new Error(`User with ID ${id} not found`);
	}

	return { user };
};
