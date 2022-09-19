import { useState, useEffect } from 'react'
import { EventsApi, Event } from '../openapi'

const EventList = () => {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [events, setEvents] = useState<Event[]>([]);
    const eventsApi = new EventsApi();

    useEffect(() => {
        eventsApi.eventsGet()
            .then(
                (result) => {
                    setIsLoaded(true);
                    const { data, status } = result
                    console.log(status)
                    setEvents(data);
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])
    if (error) {
        return <div>Error: {error}</div>;
    } else if (!isLoaded) {
        return <div>Loading...</div>;
    } else {
        return (
            <ul>
                {events.map(event => (
                    <li key={event.id}>
                        {event.title} {event.place}
                    </li>
                ))}
            </ul>
        );
    }
}

export default EventList
