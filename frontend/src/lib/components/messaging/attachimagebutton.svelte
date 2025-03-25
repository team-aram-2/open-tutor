<div class="messagesbutton mt-auto mr-0 mb-[20px] ml-[10px]">
	<!-- margin: auto 0px 20px 10px; -->

	<svg
		width="100%"
		height="100%"
		viewBox="0 0 24 24"
		fill="none"
		xmlns="http://www.w3.org/2000/svg"
		style="margin: calc(var(--font-size) * .1)"
	>
		<!-- Paperclip Path -->
		<path
			d="M21.2 11.05
				l-9.19 9.19
				a4.7 4.7 0 0 1 -7 -7
				l9.19 -9.19
				a3.2 3.2 0 1 1 4.53 4.53
				l-9.19 9.19
				a1.5 1.5 0 0 1 -2.1 -2.1
				l8.5 -8.5"
			stroke="var(--yellow-dark)"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
		/>
	</svg>
</div>

<style>
	.messagesbutton {
		display: flex;
		align-items: center;
		justify-items: center;

		width: calc(2 * var(--font-size));
		height: calc(100% - 40px);

		border-radius: var(--font-size);

		background-color: var(--yellow-light);
	}
	:hover.messagesbutton {
		transition: background-color 210ms ease-in-out;

		background-color: var(--purple-neutral);
	}
	.messagesbutton:hover svg :is(path, line, polyline) {
		transition: stroke var(--purple-light) 200ms ease-in-out;

		stroke: var(--purple-very-light);
	}
</style>
