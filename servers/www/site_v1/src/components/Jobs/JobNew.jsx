import React from "react";
import * as api from "../../api/api";

export default function JobNew(props) {

    const [job, setJob] = React.useState({
        id: "",
        name: "",
        comment: "",
        start_at: "",
    })

    function newJob() {
        api.newJob(job.job_name, job.comment).then(response => {
            setJob(response.data)
            props.created(response.data.id)
        })
    }

    return (
        <div>
            <h4>Новое задание</h4>
            <div className="mb-3">
                <label for="job_name" className="form-label">Название:</label>
                <input id="job_name" type="text" className="form-control"
                       value={job.job_name}
                       onInput={(v) => setJob({...job, job_name: v.target.value})}/>
                <label for="job_comment" className="form-label">Комментарий:</label>
                <textarea id="job_comment" type="text" className="form-control"
                          value={job.comment}
                          onInput={(v) => setJob({...job, comment: v.target.value})}/>
                <div className="d-grid gap-2 mt-3">
                    <button onClick={newJob} type="button" className="btn btn-primary">Создать</button>
                </div>
            </div>
        </div>
    )
}
