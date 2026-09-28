import { table } from "table";
import type { RsPodtype } from "../../types/configtypes.js";
import type { Node as NodeResType } from "../../types/node.js";

type ResourceMap = {
  Pod: RsPodtype[],
  Deployment: RsPodtype[],
  Namespace: RsPodtype[],
  Node: NodeResType[],
}


function getAge(timestamp: string) {
  const elapsed = Date.now() - new Date(timestamp).getTime()
  if (elapsed < 60_000) {
    return `${Math.floor(elapsed / 1000)}s`
  }

  if (elapsed < 3_600_000) {
    return `${Math.floor(elapsed / 60_000)}m`
  }

  if (elapsed < 86_400_000) {
    return `${Math.floor(elapsed / 3_600_000)}h`
  }

  return `${Math.floor(elapsed / 86_400_000)}d`

}



// FIRE: 
// i have to make this function flexable as it can be to be used for all of terminal table related work
const renderToTerminal = <T extends keyof ResourceMap>(data: ResourceMap[T], ResType: T): T | void => {

  // i could mamke the functon to pass the types of resouce is beaing send and also teh data so using if statument 
  // i can check if it's pods realted or deployment and return/render what it needs to be 
  //
  // or i could make the function more flexable and smart to handle this without senidng any kinds extra resouce type only data and still make it work
  // but i must handle the table size diffrent so that is resonabel tradeoff to push me to the first idea  
  //
  // generics cold help but not all the time 

  //
  //
  // ok final tough i wil pass the resouce type plus data and make the tableData array and map function flexable by programaticaly  



  // HOT: this neeed to be passed with tyeps like resType
  // and some data are not available and need to be programatic since some status change and dispayed in diffrent way
  // or use swich cases to help me render but it would still be huge since we have a lot of resouces 
  //
  //

  switch (ResType) {
    case "Deployment": {
      // @ts-ignore
      RenderDep(data)
      return
    }

    case "Namespace": {
      // @ts-ignore
      Namespace(data)
      return
    }

    case "Node": {
      // @ts-ignore
      NodeRes(data)
      return
    }

    case "Pod": {
      // @ts-ignore
      RenderPod(data)
      return
    }

    default: {
      console.log("dfault logs")
    }
  }
}


// HOT:
// i have to fix the sing pod rendering problem 

const RenderPod = (data: RsPodtype[]) => {
  const tableData = [
    ['NAME', 'REDY', 'STATUS', "RESTART", "AGE"],
    ...data?.map((item) => [item?.metadata?.name, `${Number(item.status.containerStatuses[0].ready)}/${item?.spec.containers?.length}`, item.metadata?.deletionTimestamp ? "terminating" : (item?.status.containerStatuses[0].state.waiting?.reason ?? item?.status?.containerStatuses[0].state?.terminated?.reason ?? item?.status?.phase), item?.status?.containerStatuses[0].restartCount,
    getAge(item?.metadata?.creationTimestamp)
    ])
  ];

  console.log(table(tableData))
}

const RenderDep = (data: RsPodtype[]) => {
}

const Namespace = (data: RsPodtype[]) => {
  const tableData = [
    ['NAME', 'STATUS', "AGE"],
    ...data?.map((item) => [item?.metadata?.name, item?.status?.phase,
    getAge(item?.metadata?.creationTimestamp)
    ])
  ];

  console.log(table(tableData))
}

const NodeRes = (data: NodeResType[]) => {

  const getNodeRoles = (item: NodeResType) => {
    const lables = item?.metadata?.labels
    let label = ""
    for (const it in lables) {
      if (it == "node-role.kubernetes.io/control-plane") {
        label = it.split("/")[1]
      }
    }

    return label
  }

  const tableData = [
    ['NAME', 'STATUS', 'ROLES', "AGE", 'VERSION'],
    ...data?.map((item) => [
      item?.metadata?.name,
      item.status?.conditions?.at(-1)?.type,
      getNodeRoles(item)?.length > 0 ? getNodeRoles(item) : "<none>",
      getAge(item?.metadata?.creationTimestamp),
      item?.status?.nodeInfo?.kubeletVersion,
    ])
  ];
  console.log(table(tableData))

}

export default renderToTerminal
