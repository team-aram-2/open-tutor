<script lang="ts">
	export let strokeColor: string = 'black';
</script>

<svg width="auto" height="100%" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
	<g transform="scale(1 -1) translate(0 -24)">
		<path d="M10 4L10 20" stroke={strokeColor} stroke-linecap="round" stroke-width="2" />
		<path
			d="M6 8L10 4L14"
			stroke={strokeColor}
			stroke-linecap="round"
			stroke-width="2"
			fill="none"
		/>
		<path
			d="M12 6L10 4L14"
			stroke={strokeColor}
			stroke-linecap="round"
			stroke-width="2"
			fill="none"
		/>
	</g>
	<text x="12.5" y="21" font-size="10" font-weight="bold" fill={strokeColor}>A</text>
	<text x="12.5" y="10.4" font-size="10" font-weight="bold" fill={strokeColor}>Z</text>
</svg>
