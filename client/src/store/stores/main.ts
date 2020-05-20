import { action, configure, observable, computed, reaction } from 'mobx';
import { Flow } from '~/domain/flows';
import {
  InteractionKind,
  Interactions,
  Link,
  Service,
  AccessPoints,
} from '~/domain/service-map';

import { StateChange } from '~/domain/misc';

import InteractionStore from './interaction';
import LayoutStore from './layout';
import RouteStore, { RouteHistorySourceKind } from './route';
import ServiceStore from './service';

import { ids } from '~/domain/ids';
import * as storage from '~/storage/local';

configure({ enforceActions: 'observed' });

export interface Props {
  historySource: RouteHistorySourceKind;
}

export class Store {
  @observable interactions: InteractionStore;

  @observable services: ServiceStore;

  @observable route: RouteStore;

  @observable layout: LayoutStore;

  @observable namespaces: Array<string> = [];

  @observable selectedTableFlow: Flow | null = null;

  constructor({ historySource }: Props) {
    this.interactions = new InteractionStore();
    this.services = new ServiceStore();
    this.route = new RouteStore(historySource);

    // LayoutStore is a store which knows geometry props of service map
    // It will be depending on flows / links as these are used to determine
    // positions of cards
    this.layout = new LayoutStore(this.services, this.interactions);

    this.restoreNamespace();
    this.setupReactions();
  }

  @action.bound
  setup({ services }: { services: Array<Service> }) {
    this.services.set(services);
  }

  @action.bound
  setNamespaces(nss: Array<string>) {
    this.namespaces = nss;

    if (!this.route.namespace && nss.length > 0) {
      this.route.goto(`/${nss[0]}`);
    }
  }

  @action.bound
  addNamespace(ns: string) {
    const nsIdx = this.namespaces.findIndex((nss: string) => nss === ns);
    if (nsIdx !== -1) return;

    this.namespaces.push(ns);
  }

  @action.bound
  removeNamespace(ns: string) {
    const nsIdx = this.namespaces.findIndex((nss: string) => nss === ns);
    if (nsIdx === -1) return;

    this.namespaces.splice(nsIdx, 1);
  }

  @action.bound
  applyNamespaceChange(ns: string, change: StateChange) {
    if (change === StateChange.Deleted) {
      this.removeNamespace(ns);
      return;
    }

    this.addNamespace(ns);
  }

  @action.bound
  selectTableFlow(flow: Flow | null) {
    this.selectedTableFlow = flow;
  }

  @action.bound
  updateInteractions<T = {}>(
    interactions: Interactions<T>,
    handleInteractions?: (kind: string, interactions: any) => void,
  ) {
    // TODO: in fact it should accurately apply a diff with current interactions

    Object.keys(interactions).forEach((k: string) => {
      const key = k as keyof Interactions<T>;

      if (key === InteractionKind.Links) {
        const links = (interactions.links || []) as Array<Link>;

        this.services.updateLinkEndpoints(links);
        this.interactions.setLinks(links);
      } else if (key === InteractionKind.Flows) {
        return;
      } else if (handleInteractions != null) {
        handleInteractions(k, interactions[key]);
      }
    });
  }

  @action.bound
  applyServiceChange(svc: Service, change: StateChange) {
    // console.log('service change: ', svc, change);

    this.services.applyServiceChange(svc, change);
  }

  @action.bound
  applyServiceLinkChange(link: Link, change: StateChange) {
    // console.log('service link change: ', link, change);

    this.interactions.applyLinkChange(link, change);
  }

  @computed
  get accessPoints(): AccessPoints {
    const index: AccessPoints = new Map();

    this.interactions.links.forEach((l: Link) => {
      const id = ids.accessPoint(l.destinationId, l.destinationPort);
      if (!index.has(l.destinationId)) {
        index.set(l.destinationId, new Map());
      }

      const serviceAPs = index.get(l.destinationId)!;
      serviceAPs.set(l.destinationPort, {
        id,
        port: l.destinationPort,
        protocol: l.ipProtocol,
        serviceId: l.destinationId,
      });
    });

    return index;
  }

  @computed
  get mocked(): boolean {
    return this.route.hash === 'mock';
  }

  private setupReactions() {
    // prettier-ignore
    reaction(() => this.route.namespace, namespace => {
      if (!namespace) return;

      storage.saveLastNamespace(namespace);
    });
  }

  private restoreNamespace() {
    if (this.route.namespace) return;

    const lastNamespace = storage.getLastNamespace();
    if (!lastNamespace) return;

    this.route.goto(`/${lastNamespace}`);
  }
}
