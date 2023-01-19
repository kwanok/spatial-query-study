import type {Component} from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import Map from "./lib/Map";

const App: Component = () => {
    return (
        <div class={styles.App}>
            <Map />
            <Map />
        </div>
    );
};

export default App;
