<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';

	enum Tab {
		Login = 'LOGIN',
		Register = 'REGISTER'
	}

	$: selectedTab = Tab.Login;
	$: submittingForm = false;
	$: submissionError = '';

	const onSubmit = async (e: any) => {
		if (submittingForm) return;

		submittingForm = true;
		submissionError = '';

		e.preventDefault();

		const formData = Object.fromEntries(new FormData(e.target));

		let response;
		switch (selectedTab) {
			case Tab.Login:
				response = fetch(`${PUBLIC_API_HOST}/login`, {
					method: 'POST',
					body: JSON.stringify(formData)
				});
				break;
			case Tab.Register:
				response = fetch(`${PUBLIC_API_HOST}/register`, {
					method: 'POST',
					body: JSON.stringify(formData)
				});
				break;
			default:
				console.error('unknown form tab');
		}

		try {
			const jsonResponse = await (await response)?.json();
		} catch (err) {
			submissionError = err as any;
		} finally {
			submittingForm = false;
		}
	};
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-1">
		<h1 class="text-center text-5xl">Open Tutor</h1>

		<div class="my-6">
			<div class="flex flex-row h-10">
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab.Login ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab.Login)}>Login</button
				>
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab.Register ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab.Register)}>Register</button
				>
			</div>
			<hr class="border-gray-600 border-dashed" />
		</div>

		<form class="flex flex-col gap-3 w-120 mx-auto" on:submit={onSubmit}>
			{#if selectedTab === Tab.Login}
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="email"
					type="email"
					placeholder="Email"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="password"
					type="password"
					placeholder="Password"
				/>
				<input
					class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50"
					type="submit"
					value="Login"
					disabled={submittingForm}
				/>
			{:else if selectedTab === Tab.Register}
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="first_name"
					type="text"
					placeholder="First name"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="last_name"
					type="text"
					placeholder="Last name"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="email"
					type="email"
					placeholder="Email"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="password"
					type="password"
					placeholder="Password"
				/>
				<input
					class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50"
					type="submit"
					value="Register"
					disabled={submittingForm}
				/>
			{/if}
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>
