exports.toBottom = (element, up = 0) => {
	const wrapper = document.createElement("div")
	element.parentNode.insertBefore(wrapper, element.nextSibling)

	return () => {
		const rect = wrapper.getBoundingClientRect()
		const size = screenSize()

		if(rect.top + rect.height > size.height) {
			const elementRect = element.getBoundingClientRect()

			element.style.position = "fixed"
			element.style.bottom = 0
			wrapper.style.height = elementRect.height + "px"
		} else {
			element.style.position = "relative"
			element.style.bottom = null
			wrapper.style.height = 0
		}
		return new Promise(resolve => { resolve() })
	};
}

function screenSize() {
	return {
		width: window.innerWidth
			|| document.documentElement.clientWidth
			|| document.body.clientWidth,
		height: window.innerHeight
			|| document.documentElement.clientHeight
			|| document.body.clientHeight
	}
}