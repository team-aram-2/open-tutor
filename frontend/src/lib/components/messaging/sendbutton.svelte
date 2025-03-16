<div class="messagesbutton mt-[20px] mr-[10px] ml-0px mb-[20px]">
	<svg
		width="100%"
		height="100%"
		viewBox="0 0 24 24"
		fill="none"
		xmlns="http://www.w3.org/2000/svg"
	>
		<!-- Vertical line -->
		<line
			x1="12"
			y1="19"
			x2="12"
			y2="5"
			stroke="var(--yellow-dark)"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
		/>
		<!-- Head of arrow -->
		<polyline
			points="5 12 12 5 19 12"
			fill="none"
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
