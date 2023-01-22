import 'leaflet/dist/leaflet.css'
import L from 'leaflet'
import {createSignal, onMount} from "solid-js";
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


const fetchLocation = async (bounds: Bounds, version: number) => {
    const URL = `http://localhost:30001/locations/v${version}/polygon?x1=${bounds.getWest()}&y1=${bounds.getSouth()}&x2=${bounds.getEast()}&y2=${bounds.getNorth()}`
    const response = await fetch(URL)
    return await response.json()
}

const buildMap = (div: HTMLDivElement, version: number, time: HTMLHeadingElement) => {
    const map = L.map(div).setView([37.5686, 126.9871], 16)

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
        maxZoom: 18,
        minZoom: 13
    }).addTo(map);

    const removeAllMarkers = () => {
        map.eachLayer((layer) => {
            if (layer instanceof L.Marker) {
                layer.remove()
            }
        })
    }

    const mapMoveEndEvent = async () => {
        const startedAt = performance.now()
        removeAllMarkers()
        const bounds = map.getBounds()
        const locations: Array<Location> = await fetchLocation(bounds, version)
        locations.forEach((location) => {
            L.marker([location.point.y, location.point.x])
                .addTo(map)
                .bindPopup(location.name)
        })

        const endedAt = performance.now()
        const elapsedTime = Math.round(endedAt - startedAt)
        time.textContent = `Elapsed Time: ${elapsedTime}ms`
    }

    map.on('moveend', mapMoveEndEvent)
}

const Map = ({title, version}: { title: string, version: number }) => {
    let mapDiv: any;
    let time: any;
    onMount(() => buildMap(mapDiv, version, time));
    return (
        <div class={styles.MapBox}>
            <h2>{title}</h2>
            <p ref={time}>Elapsed Time:</p>
            <div ref={mapDiv} id='main-map' class={styles.Map}/>
        </div>
    );
}

export default Map;
