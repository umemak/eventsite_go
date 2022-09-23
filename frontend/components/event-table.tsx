import { useState, useEffect } from 'react'
import { EventsApi, Event } from '../openapi'
import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    useReactTable,
} from '@tanstack/react-table'

const EventTable = () => {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(true);
    const [events, setEvents] = useState<Event[]>([]);
    const eventsApi = new EventsApi();

    const columnHelper = createColumnHelper<Event>()
    const columns = [
        columnHelper.accessor('id', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('title', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('start', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('place', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('open', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('close', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
        columnHelper.accessor('author', {
            cell: info => info.getValue(),
            footer: info => info.column.id,
        }),
    ]

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
    const table = useReactTable({
        data: events,
        columns: columns,
        getCoreRowModel: getCoreRowModel(),
    })

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Loading...</div>;
    } else {
        return (
            <>
                <table>
                    <thead>
                        {table.getHeaderGroups().map(headerGroup => (
                            <tr key={headerGroup.id}>
                                {headerGroup.headers.map(header => (
                                    <th key={header.id}>
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(
                                                header.column.columnDef.header,
                                                header.getContext()
                                            )}
                                    </th>
                                ))}
                            </tr>
                        ))}
                    </thead>
                    <tbody>
                        {table.getRowModel().rows.map(row => (
                            <tr key={row.id}>
                                {row.getVisibleCells().map(cell => (
                                    <td key={cell.id}>
                                        {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                    </td>
                                ))}
                            </tr>
                        ))}
                    </tbody>
                    <tfoot>
                        {table.getFooterGroups().map(footerGroup => (
                            <tr key={footerGroup.id}>
                                {footerGroup.headers.map(header => (
                                    <th key={header.id}>
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(
                                                header.column.columnDef.footer,
                                                header.getContext()
                                            )}
                                    </th>
                                ))}
                            </tr>
                        ))}
                    </tfoot>
                </table>
            </>
        );
    }
}

export default EventTable
