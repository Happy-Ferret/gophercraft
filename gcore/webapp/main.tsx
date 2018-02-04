import * as React from "react";
import * as ReactDOM from "react-dom";
import { HashRouter, Route } from "react-router-dom";

import { RegistrationPage } from "./Components/RegistrationPage";
import { AdminPanel } from "./Components/AdminPanel";

window.addEventListener("load", () => {
    ReactDOM.render(
        <HashRouter>
            <div>
                <Route exact path="/" component={RegistrationPage}></Route>
                <Route path="/admin" component={AdminPanel}></Route>
            </div>
        </HashRouter>,
        document.body
    );
});
