/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
import {Observable} from "rxjs/Observable";
import {Inject, Injectable, Injector} from "@angular/core";
import {Http} from "@angular/http";
import * as schema from "generate-schema";

import {
    IActivityContribution,
    IConnectorContribution,
    IFieldDefinition,
    IValidationResult,
    ValidationResult,
    WiContrib,
    WiContributionUtils,
    WiServiceHandlerContribution
} from "wi-studio/app/contrib/wi-contrib";

@WiContrib({})
@Injectable()
export class AppendActivityContribution extends WiServiceHandlerContribution {
    constructor(@Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    value = (fieldName: string, context: IActivityContribution): Observable<any> | any => {
        /* op is mentioned as array in activity.json, defining the type of the array here, so that it can be used in subsequent activities */
        if (fieldName === "output") {
            let inputSchema: IFieldDefinition = context.getField("inputarray");
            console.log("output schema is")
            console.log(inputSchema.value)
            console.log("-------------------")
            console.log(inputSchema.value)

            let outputSchema = schema.json(JSON.parse(inputSchema.value))
            console.log(outputSchema)
            return JSON.stringify(outputSchema,null, 2);
        }
        return null;
    };
}