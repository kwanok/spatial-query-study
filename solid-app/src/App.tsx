import type {Component} from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import Map from "./lib/Map";

const App: Component = () => {
    return (
        <div class={styles.App}>
            <div class={styles.Maps}>
                <Map title="Compare Distance" version={1}/>
                <Map title="Use Spatial Index" version={2}/>
            </div>
        </div>
    );
};

export default App;
