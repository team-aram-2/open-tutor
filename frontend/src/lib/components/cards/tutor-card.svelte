<script>
	export let profilePicture = '/images/cards/pfp.png';
	export let name = 'John Doe';
	export let rating = 2;
	export let skills = [''];
	export let userId = '';
	export let width = '';
</script>

<div class="tutor-card flex h-screen" style="width: {width}px;">
	<h1 class="tutor-header">
		<img class="pfp" src={profilePicture} alt="" />
		<a href={`/my_people/student/my_tutor/${userId}`} class="tutor-info">
			<div class="tutor-name" style="margin-top: 5px;">{name}</div>
			<div class="rating">
				{#each Array(5) as nom, index}
					<!-- TODO: TURN THIS INTO A BUTTON OR SOMETHING -->
					<svg class="star" style:display={index < rating ? 'inline' : 'none'} viewBox="0 0 51 48">
						<path d="M25,1 L31,17 L48,17 L35,29 L39,45 L25,36 L11,45 L15,29 L2,17 L19,17 Z" />
					</svg>
					<object title={nom} style="display: none;"></object>
				{/each}
			</div>
		</a>
	</h1>
	<div class="specialties">
		{#each skills as skill}
			<div class="skill-tag">{skill}</div>
		{/each}
	</div>
</div>

<style>
	.tutor-card {
		display: flex;
		flex-direction: column;

		margin: 0px 0px 0px 0px;

		border-radius: calc(var(--font-size) * 0.75);

		/* min-width: 250px; */
		height: calc(var(--font-size) * 6);

		background: #fdf8ed;
		border: 3px solid var(--yellow-neutral);
	}

	.tutor-header {
		display: flex;
		flex-direction: row;
		height: calc(55% + max(0px, (10px - (var(--font-size) * 0.125))));
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
		width: 70%;
		height: 100%;
		display: flex;
		flex-direction: column;
	}

	.specialties {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		align-items: start;

		gap: 5px;
		margin: 0px 10px 0px 20px;
		overflow: scroll;
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
</style>
