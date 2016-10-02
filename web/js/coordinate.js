const methods = {
	add: function (that) {
		return create(this.x + that.x, this.y + that.y)
	},
	subtract: function (that) {
		return create(this.x - that.x, this.y - that.y)
	},
	distance: function () {
		return Math.sqrt(this.x*this.x + this.y*this.y)
	}
}

const create = (x = 0, y = 0) => Object.create(methods, {
	x: {
		value: x,
		enumerable: true
	},
	y: {
		value: y,
		enumerable: true
	}
})

exports.methods = methods
exports.create = create

