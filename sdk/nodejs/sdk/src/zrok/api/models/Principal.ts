/* tslint:disable */
/* eslint-disable */
/**
 * zrok
 * zrok client access
 *
 * The version of the OpenAPI document: 0.3.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Principal
 */
export interface Principal {
    /**
     * 
     * @type {number}
     * @memberof Principal
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof Principal
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof Principal
     */
    token?: string;
    /**
     * 
     * @type {boolean}
     * @memberof Principal
     */
    limitless?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof Principal
     */
    admin?: boolean;
}

/**
 * Check if a given object implements the Principal interface.
 */
export function instanceOfPrincipal(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PrincipalFromJSON(json: any): Principal {
    return PrincipalFromJSONTyped(json, false);
}

export function PrincipalFromJSONTyped(json: any, ignoreDiscriminator: boolean): Principal {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'token': !exists(json, 'token') ? undefined : json['token'],
        'limitless': !exists(json, 'limitless') ? undefined : json['limitless'],
        'admin': !exists(json, 'admin') ? undefined : json['admin'],
    };
}

export function PrincipalToJSON(value?: Principal | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'email': value.email,
        'token': value.token,
        'limitless': value.limitless,
        'admin': value.admin,
    };
}
