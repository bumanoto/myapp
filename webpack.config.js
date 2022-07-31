const path = require("path");
const webpack = require("webpack");
module.exports = {
    mode: "production",
    entry: path.resolve(__dirname, 'static/js/src/index.js'),
    output: {
        path: path.resolve(__dirname, 'static/js/dist'),
        filename: "index.js"
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            }
        ]
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery',
            'window.jQuery': 'jquery'
        })
    ]
}