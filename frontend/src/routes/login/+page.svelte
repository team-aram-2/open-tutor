<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	enum Tab {
		Login = 'LOGIN',
		Register = 'REGISTER'
	}

	$: selectedTab = Tab.Login;
	$: submittingForm = false;
	$: submissionError = '';
	$: loginError = '';

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		loginError = params.get('err') ?? '';
	});

	const onSubmit = async (e: any) => {
		if (submittingForm) return;

		submittingForm = true;
		submissionError = '';

		e.preventDefault();

		const formData = Object.fromEntries(new FormData(e.target));

		let reqPromise;
		switch (selectedTab) {
			case Tab.Login:
				reqPromise = fetch(`${PUBLIC_API_HOST}/auth/login`, {
					method: 'POST',
					body: JSON.stringify(formData)
				});
				break;
			case Tab.Register:
				reqPromise = fetch(`${PUBLIC_API_HOST}/auth/register`, {
					method: 'PUT',
					body: JSON.stringify(formData)
				});
				break;
			default:
				console.error('unknown form tab');
		}

		const response = await reqPromise;
		if (!response || !response?.ok) {
			console.error(`failed to sign up/login`);
			return;
		}

		window.localStorage.setItem('SessionToken', response.headers.get('X-Session-Token')!);
		// window.location.href = '/';
	};
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-1 w-120 mx-auto">
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

		<span class="text-red-600 text-center">{loginError}</span>

		<form
			class="flex flex-col gap-3"
			method="POST"
			action="{PUBLIC_API_HOST}/auth/{selectedTab.toLowerCase()}"
		>
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
