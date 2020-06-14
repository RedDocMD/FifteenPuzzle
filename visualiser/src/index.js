import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

class Board extends React.Component {
    render() {
        return <h1>Hello</h1>
    }
}

class Tile extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <span class="tile">
                <span class="padding"></span>
                <span class="text">{this.props.value}</span>
                <span class="padding"></span>
            </span>
        )
    }
}

class App extends React.Component {
    render() {
        return (
            <div>
                <Tile value="1"/>
                <Tile value="2"/>
            </div>
        )
    }
}

ReactDOM.render(
    <App/>,
    document.getElementById('board')
)