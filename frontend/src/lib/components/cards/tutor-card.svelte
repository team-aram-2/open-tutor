<script>
	export let profilePicture = '/images/cards/pfp.png';
	export let name = 'John Doe';
	export let rating = 2;
	export let starImage = '/images/cards/star.png';
	export let emptyStarImage = '/images/cards/empty_star.png';
	export let skills = [''];
	export let userId = '';
</script>

<div class="tutor-card flex h-screen">
	<h1 class="tutor-header">
		<img class="pfp" src={profilePicture} alt="" />
		<a href={`/my_people/student/my_tutor/${userId}`} class="tutor-name no-decoration"
			>{name}
			<div class="rating">
				{#each Array(5) as nom, index}
					<!-- TODO: TURN THIS INTO A BUTTON OR SOMETHING -->
					<object title={nom}></object>
					<img class="star" src={index < rating ? starImage : emptyStarImage} alt="" />
				{/each}
			</div>
		</a>
	</h1>
	<div class="specialties">
		{#each skills as skill}
			<div class="tag">{skill}</div>
		{/each}
	</div>
</div>

<style>
	.tutor-card {
		padding-right: 20px;
		padding-bottom: 50px;
		margin-bottom: 25px;

		border-top-right-radius: 25px;
		border-bottom-right-radius: 25px;
		border-top-left-radius: 25px;
		border-bottom-left-radius: 25px;

		min-width: 500px;
		min-height: 0px;
		height: auto;
		width: 29%;
		max-height: 1000px;

		display: flex;
		flex-direction: column;

		background: #fdf8ed;
		border: 5px solid #ebad33;
	}

	.tutor-header {
		display: flex;
		flex-direction: row;
	}
	.tutor-name {
		height: 120px;
		width: auto;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;

		position: relative;
		top: 0;
		left: 0;

		padding-left: 25px;
		padding-right: 25px;
		margin: 0;

		font-size: 3rem;
		font-weight: bold;
		color: #352505;
		/* border: 2px solid blue */
	}
	.pfp {
		width: 120px;
		height: 120px;

		margin-left: 20px;

		border-radius: 15px;
		border-top-right-radius: 25px;
		border-bottom-right-radius: 25px;
		border-top-left-radius: 25px;
		border-bottom-left-radius: 25px;
		border: 2px solid #ebad33;
	}

	.rating {
		display: flex;
		flex-direction: row;
	}

	.star {
		width: 40px;
		height: 40px;
		margin-left: 5px;
	}

	.no-decoration {
		text-decoration: none;
		color: inherit;
	}

	.specialties {
		display: flex;
		flex-direction: row;
		padding-left: 10px;
	}

	.tag {
		width: auto;
		line-height: 30px;
		height: 30px;
		margin-left: 10px;
		padding: 3px 5px 3px;

		border-radius: 15px;
		border-top-right-radius: 25px;
		border-bottom-right-radius: 25px;
		border-top-left-radius: 25px;
		border-bottom-left-radius: 25px;

		font-size: 1.2rem;
		font-weight: bold;
		text-align: center;

		background: #ebad33;
		border: 3px solid #f4d390;
		color: #352505;
	}
</style>
