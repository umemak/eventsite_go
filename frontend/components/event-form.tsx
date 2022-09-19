import { useForm, SubmitHandler } from 'react-hook-form'
import { EventsApi, Event } from '../openapi'

const EventForm = () => {
    const { register, handleSubmit, formState: { errors }, } = useForm<Event>()
    const onSubmit: SubmitHandler<Event> = (data) => {
        data.author=1
        console.log(data)
        const eventsApi = new EventsApi();
        eventsApi.eventsPost(data)
    }
    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <input {...register('title', { required: true })} placeholder="タイトル" />
            <input {...register('start', { required: true })} placeholder="開催日" />
            <input {...register('place', { required: true })} placeholder="開催場所" />
            <input {...register('open', { required: true })} placeholder="開始時間" />
            <input {...register('close', { required: true })} placeholder="終了時間" />
            <input type="submit" />
        </form>
    )
}

export default EventForm
