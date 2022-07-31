const path = require("path");
const webpack = require("webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
    mode: "production",
    entry: path.resolve(__dirname, 'static/src/index.js'),
    output: {
        path: path.resolve(__dirname, 'static/dist'),
        filename: "index.js"
    },
    module: {
        rules: [
            {
                test: /\.(scss|sass|css)$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    'sass-loader'
                ]
            }
        ]
    },
    plugins: [
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery',
            'window.jQuery': 'jquery'
        }),
        new MiniCssExtractPlugin({
            // 抽出する CSS のファイル名
            filename: 'styles.css',
        }),
    ]
}