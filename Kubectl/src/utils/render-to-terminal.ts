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


const renderToTerminal = <T extends keyof ResourceMap>(data: ResourceMap[T], ResType: T): T | void => {

  switch (ResType) {
    case "Deployment": {
      // @ts-expect-error i will fix this type error when i have types for each of htem
      renderDep(data)
      return
    }

    case "Namespace": {
      // @ts-expect-error i will fix this type error when i have types for each of htem
      namespace(data)
      return
    }

    case "Node": {
      // @ts-expect-error i will fix this type error when i have types for each of htem
      nodeRes(data)
      return
    }

    case "Pod": {
      // @ts-expect-error i will fix this type error when i have types for each of htem
      renderPod(data)
      return
    }

    default: {
      console.log("dfault logs")
    }
  }
}


// HOT:
// i have to fix the sing pod rendering problem 

const renderPod = (data: RsPodtype[]) => {
  const tableData = [
    ['NAME', 'REDY', 'STATUS', "RESTART", "AGE"],
    ...(data ?? []).map((item) => [item?.metadata?.name, `${Number(item.status.containerStatuses[0].ready)}/${item?.spec.containers?.length}`, item.metadata?.deletionTimestamp ? "terminating" : (item?.status.containerStatuses[0].state.waiting?.reason ?? item?.status?.containerStatuses[0].state?.terminated?.reason ?? item?.status?.phase), item?.status?.containerStatuses[0].restartCount,
    getAge(item?.metadata?.creationTimestamp)
    ])
  ];

  console.log(table(tableData))
}

const renderDep = (data: RsPodtype[]) => {
  console.log(data)
}

const namespace = (data: RsPodtype[]) => {
  const tableData = [
    ['NAME', 'STATUS', "AGE"],
    ...(data ?? []).map((item) => [item?.metadata?.name, item?.status?.phase,
    getAge(item?.metadata?.creationTimestamp)
    ])
  ];

  console.log(table(tableData))
}

const nodeRes = (data: NodeResType[]) => {

  const getNodeRoles = (item: NodeResType) => {
    const lables = item?.metadata?.labels
    let label = ""
    for (const it in lables) {
      if (it === "node-role.kubernetes.io/control-plane") {
        label = it.split("/")[1]
      }
    }

    return label
  }

  const tableData = [
    ['NAME', 'STATUS', 'ROLES', "AGE", 'VERSION'],
    ...(data ?? []).map((item) => [
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
