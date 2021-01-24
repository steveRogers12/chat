function setViewport() {
	root.style.cssText += `; font-size: ${ Math.min(root.clientWidth, DESIGN_WIDTH) / DESIGN_WIDTH * BASE_FONT_SIZE }px;`;
}

const DESIGN_WIDTH = 375;
const BASE_FONT_SIZE = 100;
let root = document.documentElement;

export default () => {
	window.addEventListener('resize', setViewport);
	setViewport();
};