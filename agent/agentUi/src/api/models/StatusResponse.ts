/* tslint:disable */
/* eslint-disable */
/**
 * agent/agentGrpc/agent.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { AccessDetail } from './AccessDetail';
import {
    AccessDetailFromJSON,
    AccessDetailFromJSONTyped,
    AccessDetailToJSON,
    AccessDetailToJSONTyped,
} from './AccessDetail';
import type { ShareDetail } from './ShareDetail';
import {
    ShareDetailFromJSON,
    ShareDetailFromJSONTyped,
    ShareDetailToJSON,
    ShareDetailToJSONTyped,
} from './ShareDetail';

/**
 * 
 * @export
 * @interface StatusResponse
 */
export interface StatusResponse {
    /**
     * 
     * @type {Array<AccessDetail>}
     * @memberof StatusResponse
     */
    accesses?: Array<AccessDetail>;
    /**
     * 
     * @type {Array<ShareDetail>}
     * @memberof StatusResponse
     */
    shares?: Array<ShareDetail>;
}

/**
 * Check if a given object implements the StatusResponse interface.
 */
export function instanceOfStatusResponse(value: object): value is StatusResponse {
    return true;
}

export function StatusResponseFromJSON(json: any): StatusResponse {
    return StatusResponseFromJSONTyped(json, false);
}

export function StatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): StatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accesses': json['accesses'] == null ? undefined : ((json['accesses'] as Array<any>).map(AccessDetailFromJSON)),
        'shares': json['shares'] == null ? undefined : ((json['shares'] as Array<any>).map(ShareDetailFromJSON)),
    };
}

export function StatusResponseToJSON(json: any): StatusResponse {
    return StatusResponseToJSONTyped(json, false);
}

export function StatusResponseToJSONTyped(value?: StatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'accesses': value['accesses'] == null ? undefined : ((value['accesses'] as Array<any>).map(AccessDetailToJSON)),
        'shares': value['shares'] == null ? undefined : ((value['shares'] as Array<any>).map(ShareDetailToJSON)),
    };
}

