<script>
	// Any logic would go here
</script>

<main>
	<div class="container">
		<div class="titlecontainer">
			<h1 class="open">Open</h1>
			<h1 class="tutor">Tutor</h1>
		</div>
		<div class="mission-statement">
			<p>
				Our mission is to make quality education accessible to everyone by connecting students with
				expert tutors in an open learning environment.
			</p>
		</div>
	</div>
</main>

<style>
	main {
		font-family: inter;
		display: flex;
		width: 100%;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		background-color: var(--purple-very-dark);
	}

	.container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		text-align: center;
		max-width: 100%;
		padding: 2rem;
		margin: 0 auto 0 auto;
	}

	h1 {
		font-size: 6rem;
		font-weight: 700;
		margin-bottom: 2rem;
	}
	@keyframes fadeIn {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	.titlecontainer {
		display: flex;
		flex-direction: row;
		justify-content: center;
		animation: fadeIn 2s ease-in;
	}
	.open {
		color: var(--yellow-very-light);
	}
	.tutor {
		color: var(--yellow-neutral);
	}

	.mission-statement {
		width: 75%;
		font-size: 1rem;
		line-height: 1.6;
		color: var(--yellow-very-light);
		text-align: center;
		margin: 0px auto 0px auto;
	}
</style>
