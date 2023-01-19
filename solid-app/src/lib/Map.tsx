import 'leaflet/dist/leaflet.css'
import L from 'leaflet'
import {onMount} from "solid-js";
import styles from '../App.module.css';

type Bounds = {
    getNorth: () => number,
    getSouth: () => number,
    getEast: () => number,
    getWest: () => number
}

type Location = {
    id: number,
    name: string,
    point: Point
}

type Point = {
    x: number,
    y: number
}


const fetchLocation = async (bounds: Bounds) => {
    const URL = `http://localhost:30001/locations/polygon?x1=${bounds.getWest()}&y1=${bounds.getSouth()}&x2=${bounds.getEast()}&y2=${bounds.getNorth()}`
    const response = await fetch(URL)
    return await response.json()
}

const buildMap = (div: HTMLDivElement) => {
    const map = L.map(div).setView([37.5686, 126.9871], 16)

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    L.marker([37.5686, 126.9871]).addTo(map)
        .bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
        .openPopup();

    map.on('movestart', () => {
        // remove all markers
        map.eachLayer((layer) => {
            if (layer instanceof L.Marker) {
                map.removeLayer(layer)
            }
        })
    })

    map.on('moveend', () => {
        let bounds = map.getBounds()
        fetchLocation(bounds).then((locations: [Location]) => {
            locations.forEach(location => {
                L.marker([location.point.y, location.point.x]).addTo(map)
                    .bindPopup(location.name)
            })
        })
    })


}

const Map = () => {
    let mapDiv: any;
    onMount(() => buildMap(mapDiv));
    return (
        <div ref={mapDiv} id='main-map' class={styles.mapBox}/>
    );
}

export default Map;
