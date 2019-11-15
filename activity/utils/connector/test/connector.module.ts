import { NgModule } from "@angular/core";
import { TestConnectorContribution } from "./connector";
import { WiServiceContribution } from "wi-studio/app/contrib/wi-contrib";


@NgModule({
    providers: [
        {
            provide: WiServiceContribution,
            useClass: TestConnectorContribution
        }
    ]
})

export default class TestConnectorModule {

}