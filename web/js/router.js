module.exports = {
	firstA,
	getRoute,
	init,
	next,
	meetRequirements
}

function firstA(element) {
	if(element.tagName.toUpperCase() == "A") {
		return element
	}
	if(element.parentElement == null) {
		return null
	}
	return firstA(element.parentElement)
}

function next(a) {
	return getRoute(a).then(html => {
		window.history.pushState(html.body.innerHTML, window.document.title, a.href)
		const frag = document.createDocumentFragment()
		Array.from(html.body.children).forEach(c => frag.appendChild(c))
		return Promise.resolve(frag)
	}, e => {
		return Promise.reject(e)
	})
}

function getRoute(a) {
	if(a && a.tagName.toUpperCase() != "A" && a.href) {
		return Promise.reject(new Error("Provided value is not an A element with an href"))
	}

	return Bliss.fetch(a.href + "?theme=none", { responseType: "document"}).then(response => {
		return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error("Did not get a document back from " + a.href))
	}, e => {
		return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error("Did not get a document back from " + a.href))
	})
}

function init() {
	const page = window.location.origin + window.location.pathname
	window.history.replaceState(Bliss("main").innerHTML, window.document.title, page)
}

function meetRequirements() {
	return !!(
		DOMParser &&
		window.history &&
		window.history.pushState &&
		document.implementation &&
		document.implementation.createHTMLDocument
	)
}