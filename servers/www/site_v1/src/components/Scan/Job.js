import React from "react";

export default function Job(props) {
    return <div className={"alert alert-success"} role={"alert"}>
        <h4 className={"alert-heading"}>{props.job.name}</h4>
        id: {props.job.id}
        <br/>
        start: {props.job.start_at}
        <br/>
        Comment: {props.job.comment}
    </div>
}
