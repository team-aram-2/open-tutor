<script>
	export let profilePicture = '/images/cards/pfp.png';
	export let name = 'john doe';
	export let rating = 2;
	export let skills = [''];
	// export let userId = '';
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

		border-top-right-radius: 25px;
		border-bottom-right-radius: 25px;
		border-top-left-radius: 25px;
		border-bottom-left-radius: 25px;

		display: flex;
		flex-direction: column;

		background: var(--yellow-very-light);
		border: 5px solid var(--yellow-neutral);
	}

	.tutor-header {
		display: flex;
		flex-direction: row;
		/* border: 2px solid blue; */
		max-height: 9rem;
		height: 9em;
	}
	.tutor-name {
		height: 4.5rem;
		width: 90%;
		max-width: 90%;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;

		position: relative;
		top: 0;
		left: 0;

		padding-left: 15px;
		padding-right: 15px;
		margin: 0;

		font-size: 3rem;
		font-weight: bold;
		color: var(--yellow-very-dark);
	}

	.pfp {
		width: 120px;
		height: 120px;

		margin-left: 20px;
		margin: 15px 0 20px 20px;

		border-radius: 25px;
		border: 3px solid var(--yellow-neutral);
	}

	.rating {
		/* border: 2px solid chartreuse; */
		display: flex;
		flex-direction: row;
		overflow: visible;
		height: 4em;
		width: 100%;
		margin-left: 10px;
		bottom: 0;
		justify-content: flex-start;
	}

	.star {
		font-size: inherit;
		max-width: 20%;
		width: 3.75em;
		height: 3.75em;
		max-height: 3.75em;
		margin-right: 5px;
		fill: var(--yellow-neutral);
		overflow: visible;
	}

	.tutor-info {
		text-decoration: none;
		color: inherit;
		width: 70%;
		height: 100%;
		display: flex;
		flex-direction: column;
	}

	.specialties {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		padding-top: 3px;
		height: 50%;
		margin: 5px 10px 5px 20px;
	}

	.skill-tag {
		display: flex;
		align-items: center;
		justify-content: center;

		width: auto;
		height: 1.75em;
		margin: 0 5px 7px 0;
		padding: 3px 5px 3px 5px;

		border-radius: 25px;

		font-size: 1.2rem;
		font-weight: bold;
		text-align: center;

		background: var(--yellow-light);
		border: 3px solid var(--yellow-neutral);
		white-space: nowrap;
	}
</style>
