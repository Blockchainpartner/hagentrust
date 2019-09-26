import { BigDecimal, BigInt, log, store } from "@graphprotocol/graph-ts"
import {
  Contract,
  ClaimSet,
  ClaimRemoved
} from "../generated/Contract/Contract"
import { ExampleEntity } from "../generated/schema"

export function handleClaimSet(event: ClaimSet): void {
  // Entities can be loaded from the store using a string ID; this ID
  // needs to be unique across all entities of the same type
  let ID = event.params.issuer.toHex()+event.params.recipient.toHex()+event.params.topic.toHex();


  let entity = ExampleEntity.load(ID);

  // Entities only exist after they have been saved to the store;
  // `null` checks allow to create entities on demand
  if (entity == null) {
    entity = new ExampleEntity(ID);
  }

  // Entity fields can be set based on event parameters
  entity.issuer = event.params.issuer;
  entity.recipient = event.params.recipient;
  entity.topic = event.params.topic;

  let gasUsed = event.block.gasUsed;
  let gasLimit = event.block.gasLimit;
  let gasLimitDecimal = new BigDecimal(gasLimit);
  let gasUsedDecimal = new BigDecimal(gasUsed);
  let ratio = gasUsedDecimal.div(gasLimitDecimal);

  entity.trustValue = ratio;

  // Entities can be written to the store with `.save()`
  entity.save()

  // Note: If a handler doesn't require existing field values, it is faster
  // _not_ to load the entity from the store. Instead, create it fresh with
  // `new Entity(...)`, set the fields that should be updated and save the
  // entity back to the store. Fields that were not set or unset remain
  // unchanged, allowing for partial updates to be applied.

  // It is also possible to access smart contracts from mappings. For
  // example, the contract that has emitted the event can be connected to
  // with:
  //
  // let contract = Contract.bind(event.address)
  //
  // The following functions can then be called on this contract to access
  // state variables and other data:
  //
  // - contract.hasIssuedTopic(...)
  // - contract.hasIssuedForRecipient(...)
  // - contract.getIssuerForRecipientTopic(...)
  // - contract.getIssuerForRecipient(...)
  // - contract.getIssuerTopicForRecipient(...)
  // - contract.isRecipient(...)
  // - contract.getIssuers(...)
  // - contract.getRecipientForIssuerTopic(...)
  // - contract.getRecipients(...)
  // - contract.getClaim(...)
  // - contract.registry(...)
  // - contract.getTopics(...)
}

export function handleClaimRemoved(event: ClaimRemoved): void {
  store.remove('ExampleEntity', event.params.issuer.toHex()+event.params.recipient.toHex()+event.params.topic.toHex())
}
