// export function resizeObserver(node: HTMLElement, callback: (width: number) => void) {
// 	console.log("starting resizeobserver");
// 	const observer = new ResizeObserver((entries) => {
// 		for (const entry of entries) {
// 			callback(entry.contentRect.width);
// 		}
// 	});
// 	observer.observe(node);
// 	console.log("ending resizeobserver");

// 	return {
// 		destroy() {
// 			observer.disconnect();
// 		}
// 	};
// }
