/* tslint:disable */
/* eslint-disable */
/**
 * zrok
 * zrok client access
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { SparkDataSample } from './SparkDataSample';
import {
    SparkDataSampleFromJSON,
    SparkDataSampleFromJSONTyped,
    SparkDataSampleToJSON,
    SparkDataSampleToJSONTyped,
} from './SparkDataSample';

/**
 * 
 * @export
 * @interface Environment
 */
export interface Environment {
    /**
     * 
     * @type {string}
     * @memberof Environment
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof Environment
     */
    host?: string;
    /**
     * 
     * @type {string}
     * @memberof Environment
     */
    address?: string;
    /**
     * 
     * @type {string}
     * @memberof Environment
     */
    zId?: string;
    /**
     * 
     * @type {boolean}
     * @memberof Environment
     */
    remoteAgent?: boolean;
    /**
     * 
     * @type {Array<SparkDataSample>}
     * @memberof Environment
     */
    activity?: Array<SparkDataSample>;
    /**
     * 
     * @type {boolean}
     * @memberof Environment
     */
    limited?: boolean;
    /**
     * 
     * @type {number}
     * @memberof Environment
     */
    createdAt?: number;
    /**
     * 
     * @type {number}
     * @memberof Environment
     */
    updatedAt?: number;
}

/**
 * Check if a given object implements the Environment interface.
 */
export function instanceOfEnvironment(value: object): value is Environment {
    return true;
}

export function EnvironmentFromJSON(json: any): Environment {
    return EnvironmentFromJSONTyped(json, false);
}

export function EnvironmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): Environment {
    if (json == null) {
        return json;
    }
    return {
        
        'description': json['description'] == null ? undefined : json['description'],
        'host': json['host'] == null ? undefined : json['host'],
        'address': json['address'] == null ? undefined : json['address'],
        'zId': json['zId'] == null ? undefined : json['zId'],
        'remoteAgent': json['remoteAgent'] == null ? undefined : json['remoteAgent'],
        'activity': json['activity'] == null ? undefined : ((json['activity'] as Array<any>).map(SparkDataSampleFromJSON)),
        'limited': json['limited'] == null ? undefined : json['limited'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}

export function EnvironmentToJSON(json: any): Environment {
    return EnvironmentToJSONTyped(json, false);
}

export function EnvironmentToJSONTyped(value?: Environment | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'description': value['description'],
        'host': value['host'],
        'address': value['address'],
        'zId': value['zId'],
        'remoteAgent': value['remoteAgent'],
        'activity': value['activity'] == null ? undefined : ((value['activity'] as Array<any>).map(SparkDataSampleToJSON)),
        'limited': value['limited'],
        'createdAt': value['createdAt'],
        'updatedAt': value['updatedAt'],
    };
}

