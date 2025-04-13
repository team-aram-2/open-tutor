<script>
	export let profilePicture = '/images/cards/pfp.png';
	export let name = 'john doe';
	export let rating = 2;
	export let skills = [''];
	export let userId = '';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';
	import { goto } from '$app/navigation';

	const createConversation = async () => {
		try {
			const sendArray = [$user_id, userId];
			const res = await fetch(`${PUBLIC_API_HOST}/conversation`, {
				method: 'POST',
				credentials: 'include',
				body: JSON.stringify(sendArray)
			});
			if (!res.ok) {
				throw new Error(`API Error: ${res.status} ${res.statusText}`);
			}
			await res.json();
			goto('/messages/student');
		} catch (err) {
			console.log('Creating convo err' + err);
		}
	};
</script>

<div class="tutor-card flex h-screen">
	<h1 class="tutor-header">
		<img class="pfp" src={profilePicture} alt="" />
		<div class="tutor-info">
			<h2 class="tutor-name no-decoration">{name}</h2>
			<div class="rating">
				{#each Array(5) as nom, index}
					<!-- TODO: TURN OBJECT TAG INTO SOMETHING ELSE -->
					<object title={nom}></object>
					<svg class="star" style:display={index < rating ? 'inline' : 'none'} viewBox="0 0 51 48">
						<path d="M25,1 L31,17 L48,17 L35,29 L39,45 L25,36 L11,45 L15,29 L2,17 L19,17 Z" />
					</svg>
				{/each}
			</div>
		</div>
		<button class="message-btn" on:click={createConversation}>
			<svg
				class="message-icon"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
			</svg>
		</button>
	</h1>
	<div class="specialties">
		{#each skills as skill}
			<div class="skill-tag">{skill}</div>
		{/each}
	</div>
</div>

<style>
	.tutor-card {
		height: 100vh;
		width: 100%;

		padding-right: 20px;
		padding-bottom: 50px;
		margin-bottom: 25px;

		border-radius: 25px;

		display: flex;
		flex-direction: column;

		background: var(--yellow-very-light);
		border: 5px solid var(--yellow-neutral);
	}

	.tutor-header {
		display: flex;
		flex-direction: row;
		/* border: 2px solid blue; */
		height: calc(var(--font-size) * 3.5);
	}
	.tutor-name {
		height: calc(var(--font-size) * 1.5);
		width: 90%;
		max-width: 90%;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
		line-height: 110%;

		position: relative;
		top: 0;
		left: 0;

		padding-left: 15px;
		padding-right: 15px;
		margin: 0;

		font-size: calc(var(--font-size) * 1.3);
		font-weight: bold;
		color: var(--yellow-very-dark);
	}

	.pfp {
		width: calc(var(--font-size) * 2.7);
		height: calc(var(--font-size) * 2.7);

		margin-left: 20px;
		margin: calc(var(--font-size) * 0.25) 0 20px 20px;

		border-radius: calc(var(--font-size) * 0.25);
		border: 3px solid var(--yellow-neutral);
	}

	.rating {
		display: flex;
		justify-content: flex-start;
		flex-direction: row;
		overflow: visible;

		height: calc(var(--font-size) * 0.5);
		width: 100%;
		margin-left: 10px;
		bottom: 0;

		background-color: transparent;
	}

	.star {
		font-size: inherit;
		max-width: 20%;
		width: auto;
		height: calc(var(--font-size) * 1.25);
		margin-right: 5px;
		fill: var(--yellow-neutral);
		overflow: visible;
	}

	.tutor-info {
		text-decoration: none;
		color: inherit;
		width: 95%;
		height: 100%;
		display: flex;
		flex-direction: column;
	}

	.specialties {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		flex: 1;
		align-items: start;

		padding-top: 3px;
		height: 50%;
		margin: 5px 10px 5px 20px;
	}

	.skill-tag {
		display: flex;
		align-items: center;
		justify-content: center;

		width: auto;
		height: calc(var(--font-size));
		margin: 0 5px 7px 0;
		padding: 3px calc(var(--font-size) * 0.25) 3px calc(var(--font-size) * 0.25);

		border-radius: calc(var(--font-size) * 0.5);

		font-size: calc(var(--font-size) * 0.6);
		font-weight: bold;
		text-align: center;

		background: var(--yellow-light);
		border: 3px solid var(--yellow-neutral);
		white-space: nowrap;
	}
	.message-btn {
		display: flex;
		align-items: center;
		justify-content: center;
		height: calc(var(--font-size) * 1.8);
		padding: 0 calc(var(--font-size) * 0.5);
		margin-right: 0px;
		border-radius: calc(var(--font-size) * 0.5);
		background: var(--yellow-neutral);
		color: var(--yellow-very-dark);
		font-weight: bold;
		font-size: calc(var(--font-size) * 0.8);
		border: none;
		cursor: pointer;
		transition: background-color 0.2s;
	}
	.message-btn:hover {
		background: var(--yellow-dark);
	}
	.message-icon {
		width: calc(var(--font-size) * 1);
		height: calc(var(--font-size) * 1);
	}
</style>
