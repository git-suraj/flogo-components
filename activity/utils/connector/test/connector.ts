import { Observable } from "rxjs/Observable";
import { Injectable, Injector, Inject } from "@angular/core";
import { Http, Response, Headers } from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    IActivityContribution,
    IConnectorContribution,
    ActionResult,
    IActionResult,
    AUTHENTICATION_TYPE
} from "wi-studio/app/contrib/wi-contrib";


@WiContrib({})

@Injectable()
/**
 *  function variables
 *  value : (fieldName: string, context: IActivityContribution) => Observable<any> | any
 *  validate : (fieldName: string, context: IActivityContribution)=> Observable<IValidationResult> | IValidationResult
 *  action? : (actionId: string, context: IActivityContribution)=> Observable<IActionResult> | IActionResult
 */
export class TestConnectorContribution extends WiServiceHandlerContribution {
    constructor(@Inject(Injector) injector) {
        super(injector);
    }
    /**
     * value is a function Variable which can be implemented inline as shown below
     */
    value = (fieldName: string, context: IConnectorContribution): Observable<any> | any => {
        console.log("******************* >>>> " + fieldName)
        return null;
    }
    /**
     * validate is a function Variable which can be implemented inline as shown below
     */
    validate = (fieldName: string, context: IConnectorContribution): Observable<IValidationResult> | IValidationResult => {
        console.log("=========== >>>> " + fieldName)
        if (fieldName === "Create" ||fieldName === "Test Connection") {
            let name: IFieldDefinition = context.getField("name");
            let region: IFieldDefinition = context.getField("region");
            let username: IFieldDefinition = context.getField("username");
            let password: IFieldDefinition = context.getField("password");
            if (name.value && region.value && username.value && password.value) {
                return ValidationResult.newValidationResult().setReadOnly(false);
            } 

        }
        return null;    
    }
    /**
     * handleAction is an optional function Variable which can be implemented inline as shown below.
     * It is optional because most UI does not need form button actions except in Connectors
     */
    action = (actionId: string, context: IConnectorContribution): Observable<IActionResult> | IActionResult => {
        console.log("******** ActionID: "+ actionId);
        if (actionId == "Test Connection") {
            console.log("***** TEST **** ")
            let name: IFieldDefinition = context.getField("name");
            let region: IFieldDefinition = context.getField("region");
            let username: IFieldDefinition = context.getField("username");
            let password: IFieldDefinition = context.getField("password");
            return ActionResult.newActionResult().setSuccess(true).setResult("Missing implementation");
        }

        if (actionId === 'Create') {
            return Observable.create(observer => {
                let actionResult = {
                    context: context,
                    authType: AUTHENTICATION_TYPE.BASIC,
                    authData: {}
                };
                observer.next(ActionResult.newActionResult().setSuccess(true).setResult(actionResult));
            });
        }
    }
}
