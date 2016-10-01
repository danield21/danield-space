function Coordinate(x=0, y=0) {
    this.x = x
    this.y = y
}

Coordinate.prototype.add = function (that) {
    return new Coordinate(this.x + that.x, this.y + that.y)
}

Coordinate.prototype.subtract = function (that) {
    return new Coordinate(this.x - that.x, this.y - that.y)
}

Coordinate.prototype.distance = function () {
    return Math.sqrt(this.x*this.x + this.y*this.y)
}

Coordinate.prototype.toString = function () {
    return `(${this.x}, ${this.y})`
}

Coordinate.create = function (...args) {
    const inst = Object.create(Coordinate.prototype)
    Coordinate.apply(inst, args)
    return inst;
}

Coordinate.assign = function (obj, ...args) {
    Object.assign(obj, Coordinate.prototype)
    Coordinate.apply(obj, args)
    return obj;
}

Coordinate.from = function (obj) {
    return Object.assign({}, obj, Coordinate.prototype)
}

module.exports = Coordinate
