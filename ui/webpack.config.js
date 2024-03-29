const path = require('path');
const VueLoaderPlugin = require('vue-loader/lib/plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: {
        bundle: ['./src/index.js'],
        bookmarklet: ['./src/bookmarklet.js'],
    },
    output: {
        filename: './[name].js',
    },
    devServer: {
        port: 8080,
        host: 'localhost',
        proxy: {
            "/v1": {
                target: "http://localhost:8080",
                router: () => 'http://localhost:3000',
                logLevel: 'debug'
            }
        }
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js'
        },
        extensions: ['*', '.js', '.vue', '.json']
    },
    module: {
        rules: [
            {
                test: /\.(txt)$/,
                loader: 'file-loader',
                options: {
                    name: '[name].[ext]',
                },
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
            {
                test: /\.(json)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(scss)$/,
                use: [
                    {
                        loader: 'style-loader', // inject CSS to page
                    },
                    {
                        loader: 'css-loader', // translates CSS into CommonJS modules
                    },
                    {
                        loader: 'sass-loader' // compiles Sass to CSS
                    }
                ]
            },
        ]
    },
    plugins: [
        new VueLoaderPlugin(),
        new HtmlWebpackPlugin({
            title: 'SHRLS Admin',
            chunks: ['bundle'],
        }),
    ]
};
