import {id as pluginId} from './manifest';

import PostType from './components/post_type';

export default class DemoPlugin {
    initialize(registry) {
        registry.registerPostTypeComponent('custom_test_plugin', PostType);
    }

    uninitialize() {
        //eslint-disable-next-line no-console
        console.log(pluginId + '::uninitialize()');
    }
}
