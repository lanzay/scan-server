import React from "react";

export default function JobList(props) {

    if (!props.jobs) return null

    let list = props.jobs.map((job, i) => {
        return (
            <div className="accordion-item row">
                <h2 className="accordion-header" id={`panelsStayOpen-headingOne-${job.id}`}>
                    <button className="accordion-button" type="button" data-bs-toggle="collapse"
                            data-bs-target={`#panelsStayOpen-collapseOne-${job.id}`} aria-expanded="false"
                            aria-controls={`panelsStayOpen-collapseOne-${job.id}`}>
                        {job.name} ({job.id})
                    </button>
                </h2>
                <div id={`panelsStayOpen-collapseOne-${job.id}`} className="accordion-collapse collapse"
                     aria-labelledby={`panelsStayOpen-headingOne${job.id}`}>
                    <div className="accordion-body">
                        <div className="row">
                            <label htmlFor="id" className="col-sm-2 col-form-label">ID</label>
                            <div className="col-sm-10">
                                <input type="text" readOnly className="form-control" id="id" value={job.id}/>
                            </div>
                        </div>
                        <div className="row">
                            <label htmlFor="name" className="col-sm-2 col-form-label">Name</label>
                            <div className="col-sm-10">
                                <input type="text" readOnly className="form-control" id="name" value={job.name}/>
                            </div>
                        </div>

                        <div className="row">
                            <label htmlFor="comment" className="col-sm-2 col-form-label">Comment</label>
                            <div className="col-sm-10">
                                <input type="text" readOnly className="form-control" id="comment" value={job.comment}/>
                            </div>
                        </div>
                        <div className="row">
                            <label htmlFor="start_at" className="col-sm-2 col-form-label">Start</label>
                            <div className="col-sm-10">
                                <input type="text" readOnly className="form-control" id="start_at"
                                       value={job.start_at}/>
                            </div>
                        </div>

                        <div className="row">
                            <button id={job.id} onClick={props.onClickJobOpen} type={"button"}
                                    className={"btn btn-primary btn-lg m-2"}>Открыть
                            </button>
                            <button id={job.id} onClick={props.onClickJobCloce} type={"button"}
                                    className={"btn btn-warning btn-sm m-2"}>Закрыть
                            </button>
                        </div>

                    </div>
                </div>
            </div>

        )
    })

    return <div className="accordion accordion-flush" id="list-job">
        {list}
    </div>

}
