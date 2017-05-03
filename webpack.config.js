/* eslint-env node */
var path = require('path'),
    WebpackCopy = require('copy-webpack-plugin'),
    ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: {
        app: './web/js/app.js',
        styles: './web/sass/app.scss'
    },
    output: {
        path: path.join(__dirname, 'app', 'dist'),
        filename: 'js/[name].js',
        library: '[name]'
    },
    module: {
        rules: [{
            test: /\.modernizrrc(\.json)?$/,
            use: [
                'modernizr-loader',
                'json-loader'
            ]
        }, {
            test: /\.js$/,
            use: [{
                loader: 'babel-loader',
                options: {
                    presets: ['es2015']
                }
            }]
        }, {
            test: /\.scss$/,
            use: ExtractTextPlugin.extract({
                fallback: 'style-loader',
                use: [
                    'css-loader',
                    'sass-loader'
                ]
            })
        }]
    },
    plugins: [
        new WebpackCopy([
            { from: path.join('web', 'view'), to: path.join('..', 'view') },
            { from: path.join('web', 'svg'), to: path.join('svg') },
            { from: path.join('web', 'lib'), to: path.join('js') }
        ]),
        new ExtractTextPlugin('css/[name].css')
    ],
    resolve: {
        alias: {
            modernizr$: path.resolve(__dirname, 'web/js/.modernizrrc')
        }
    }
}