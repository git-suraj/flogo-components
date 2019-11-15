/**
 * Imports
 */
import { Observable } from "rxjs/Observable";
import { Injectable, Injector, Inject } from "@angular/core";
import { Http } from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    IActivityContribution,
    IConnectorContribution,
    WiContributionUtils
} from "wi-studio/app/contrib/wi-contrib";

/**
 * Main
 */
@WiContrib({})
@Injectable()

export class TestConnectionActivityUIContributionHandler extends WiServiceHandlerContribution {
    constructor( @Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    /**
     * The value object allows you to specify what types of values you can pick for a certain field
     */
    value = (fieldName: string, context: IActivityContribution): Observable<any> | any => {
        /**
         * For the field iftttConnection the only allowed types are connections that
         * are created as an iftttConnector (the connector category as specified in the 
         * connector.json must match what we specify here)
         */
        if (fieldName === "testConnection") {
            //  return WiContributionUtils.getConnections("tibco-sqs");
            return Observable.create(observer => {
                let connectionRefs = [];
                /**
                 * The category is IFTTT
                 */
                WiContributionUtils.getConnections(this.http, "utils").subscribe((data: IConnectorContribution[]) => {
                    data.forEach(connection => {
                        console.log("--------------- <<>>> ")
                        console.log(connection)
                        for (let i = 0; i < connection.settings.length; i++) {
                            if (connection.settings[i].name === "name") {
                                connectionRefs.push({
                                    "unique_id": WiContributionUtils.getUniqueId(connection),
                                    "name": connection.settings[i].value
                                });
                                break;
                            }
                        }
                    });
                    observer.next(connectionRefs);
                });
            });
        } else {
            return null;
        }
    }

    /**
     * The validate object can be used to validate the input of certain fields
     */
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        /**
         * For the field iftttConnection check that the connection has been set, otherwise
         * display the errormessage
         */
        if (fieldName === "testConnection") {
            let connection: IFieldDefinition = context.getField("testConnection")
            if (connection.value === null) {
                return ValidationResult.newValidationResult().setError("AUTHENTICATE-1000", "Test Connection must be configured");
            }
        }
        return null;
    }
}
