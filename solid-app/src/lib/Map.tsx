import 'leaflet/dist/leaflet.css'
import L from 'leaflet'
import { onMount} from "solid-js";
import styles from '../App.module.css';

const fetchLocation = async () => {
    (await fetch(`http://localhost:8080`)).json().then(console.log)
}
const buildMap = (div: HTMLDivElement) => {
    const map = L.map(div).setView([37.5686, 126.9871], 16)

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    L.marker([37.5686, 126.9871]).addTo(map)
        .bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
        .openPopup();
}

const Map = () => {
    let mapDiv: any;
    onMount(() => buildMap(mapDiv));
    return (
        <div ref={mapDiv} id='main-map' class={styles.mapBox}/>
    );
}

export default Map;
